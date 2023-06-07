use reqwest::Error;
use serde::de::DeserializeOwned;
use serde::{de, Deserialize, Deserializer, Serialize};
use serde_json::Value;
use std::collections::HashMap;

// https://play.rust-lang.org/?version=stable&mode=debug&edition=2018&gist=ee7f582b5873013723596790a7993925
fn de_price<'de, D: Deserializer<'de>>(deserializer: D) -> Result<f64, D::Error> {
    Ok(match Value::deserialize(deserializer)? {
        Value::String(s) => s.parse().map_err(de::Error::custom)?,
        Value::Number(num) => num.as_f64().ok_or(de::Error::custom("Invalid number"))? as f64,
        _ => return Err(de::Error::custom("wrong type")),
    })
}

#[derive(Serialize, Deserialize, Debug)]
struct Price {
    #[serde(deserialize_with = "de_price")]
    bid: f64,
    #[serde(deserialize_with = "de_price")]
    ask: f64,
    #[serde(deserialize_with = "de_price")]
    last: f64,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct LatestPrices {
    status: String,
    message: Option<String>,
    prices: HashMap<String, Price>,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct LatestPriceForCoin {
    status: String,
    message: Option<String>,
    prices: Price,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct LatestPrice {
    status: String,
    message: Option<String>,
    rate: u64,
    market: String,
}

#[derive(Debug)]
pub enum PriceResult {
    LatestPrices(LatestPrices),
    LatestPriceForCoin(LatestPriceForCoin),
}

async fn get<T>(url: &str) -> Result<T, Error>
where
    T: DeserializeOwned,
{
    let resp = reqwest::get(url).await?.json().await?;
    Ok(resp)
}

pub async fn get_latest_prices(
    cointype: Option<String>,
    markettype: Option<String>,
) -> Result<PriceResult, Error> {
    match (cointype, markettype) {
        (None, None) => get::<LatestPrices>("https://www.coinspot.com.au/pubapi/v2/latest")
        .await.map(PriceResult::LatestPrices),
        (None, Some(_)) => todo!(),
        (Some(coin), None) => get::<LatestPriceForCoin>(&format!("https://www.coinspot.com.au/pubapi/v2/latest/{coin}"))
        .await.map(PriceResult::LatestPriceForCoin),
        (Some(coin), Some(market)) => get::<LatestPriceForCoin>(&format!("https://www.coinspot.com.au/pubapi/v2/latest/{coin}/{market}"))
        .await.map(PriceResult::LatestPriceForCoin),
    }
}

pub async fn get_latest_buy_price(
    cointype: String,
    markettype: Option<String>,
) -> Result<LatestPrice, Error> {
    let url = if let Some(market) = markettype {
        format!("https://www.coinspot.com.au/pubapi/v2/buyprice/{cointype}/{market}/")
    } else {
        format!("https://www.coinspot.com.au/pubapi/v2/buyprice/{cointype}")
    };
    get::<LatestPrice>(&url).await
}

pub async fn get_latest_sell_price(
    cointype: String,
    markettype: Option<String>,
) -> Result<LatestPrice, Error> {
    let url = if let Some(market) = markettype {
        format!("https://www.coinspot.com.au/pubapi/v2/sellprice/{cointype}/{market}/")
    } else {
        format!("https://www.coinspot.com.au/pubapi/v2/sellprice/{cointype}")
    };
    get::<LatestPrice>(&url).await
}