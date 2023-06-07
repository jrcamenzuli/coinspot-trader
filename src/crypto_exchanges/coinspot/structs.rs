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

pub enum TransactionType {
    BUY,
    SELL,
}

pub enum OrderType {
    OPEN,
    COMPLETED,
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
    #[serde(deserialize_with = "de_price")]
    rate: f64,
    market: String,
}

#[derive(Serialize, Deserialize, Debug)]
struct Order {
    amount: f64,
    rate: f64,
    total: f64,
    coin: String,
    market: String,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct Orders {
    status: String,
    message: Option<String>,
    buyorders: Vec<Order>,
    sellorders: Vec<Order>,
}