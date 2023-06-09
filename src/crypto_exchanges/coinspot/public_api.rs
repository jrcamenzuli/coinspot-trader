use super::structs::{LatestPrice, LatestPrices, LatestPriceForCoin, TransactionType, TransactionType::{BUY,SELL,SWAP}, Orders, OrderType};
use reqwest::Error;
use super::get;

const BASE_URL: &str = "https://www.coinspot.com.au/pubapi/v2/";

pub async fn get_latest_prices() -> Result<LatestPrices, Error> {
    get(&format!("{BASE_URL}latest")).await
}

pub async fn get_latest_price(
    coin_type: String,
    market_type: Option<String>,
) -> Result<LatestPriceForCoin, Error> {
    let url = match market_type {
        None => format!("{BASE_URL}latest/{coin_type}"),
        Some(market) => format!("{BASE_URL}latest/{coin_type}/{market}"),
    };
    get(&url).await
}

pub async fn get_latest_transaction_price(
    coin_type: String,
    market_type: Option<String>,
    transaction_type: Option<TransactionType>,
) -> Result<LatestPrice, Error> {
    let url = match (market_type, transaction_type) {
        (None, None) => todo!(),
        (Some(_), None) => todo!(),
        (None, Some(BUY)) => format!("{BASE_URL}buyprice/{coin_type}"),
        (None, Some(SELL)) => format!("{BASE_URL}sellprice/{coin_type}"),
        (Some(market), Some(BUY)) => format!("{BASE_URL}buyprice/{coin_type}/{market}"),
        (Some(market), Some(SELL)) => format!("{BASE_URL}sellprice/{coin_type}/{market}"),
        (_,Some(SWAP)) => todo!()
    };
    get(&url).await
}

pub async fn get_orders(
    coin_type: String,
    order_type: OrderType,
    market_type: Option<String>,
) -> Result<Orders, Error> {
    let url = match (order_type, market_type) {
        (OrderType::OPEN, None) => format!("{BASE_URL}orders/open/{coin_type}"),
        (OrderType::OPEN, Some(market)) => format!("{BASE_URL}orders/open/{coin_type}/{market}"),
        (OrderType::COMPLETED, None) => format!("{BASE_URL}orders/completed/{coin_type}"),
        (OrderType::COMPLETED, Some(market)) => format!("{BASE_URL}orders/completed/{coin_type}/{market}"),
    };
    get(&url).await
}