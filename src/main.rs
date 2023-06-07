use reqwest::Error;
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
struct LatestPrices {
    status: String,
    message: Option<String>,
    prices: HashMap<String, Price>,
}

async fn get_latest_prices() -> Result<LatestPrices, Error> {
    let resp = reqwest::get("https://www.coinspot.com.au/pubapi/v2/latest")
        .await?
        .json::<LatestPrices>()
        .await?;
    Ok(resp)
}

#[tokio::main]
async fn main() {
    match get_latest_prices().await {
        Ok(prices) => {
            // Handle successful response
            println!("Latest prices: {:#?}", prices);
        }
        Err(err) => {
            // Handle error
            eprintln!("Error fetching latest prices: {}", err);
        }
    }
}
