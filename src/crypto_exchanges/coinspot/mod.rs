pub mod structs;

use reqwest::Error;
use serde::de::DeserializeOwned;
use structs::{LatestPrice, LatestPrices, LatestPriceForCoin, TransactionType};

const BASE_URL: &str = "https://www.coinspot.com.au/pubapi/v2/";

async fn get<T>(url: &str) -> Result<T, Error>
where
    T: DeserializeOwned,
{
    let resp = reqwest::get(url).await?.json().await?;
    Ok(resp)
}

pub async fn get_latest_prices() -> Result<LatestPrices, Error> {
    get::<LatestPrices>(&format!("{BASE_URL}latest")).await
}

pub async fn get_latest_price(
    coin_type: String,
    market_type: Option<String>,
) -> Result<LatestPriceForCoin, Error> {
    let url = match market_type {
        None => format!("{BASE_URL}latest/{coin_type}"),
        Some(market) => format!("{BASE_URL}latest/{coin_type}/{market}"),
    };
    get::<LatestPriceForCoin>(&url).await
}

pub async fn get_latest_transaction_price(
    coin_type: String,
    market_type: Option<String>,
    transaction_type: Option<TransactionType>,
) -> Result<LatestPrice, Error> {
    let url = match (market_type, transaction_type) {
        (None, None) => todo!(),
        (Some(_), None) => todo!(),
        (None, Some(TransactionType::BUY)) => format!("{BASE_URL}buyprice/{coin_type}"),
        (None, Some(TransactionType::SELL)) => format!("{BASE_URL}sellprice/{coin_type}"),
        (Some(market), Some(TransactionType::BUY)) => format!("{BASE_URL}buyprice/{coin_type}/{market}"),
        (Some(market), Some(TransactionType::SELL)) => format!("{BASE_URL}sellprice/{coin_type}/{market}"),
    };
    get::<LatestPrice>(&url).await
}