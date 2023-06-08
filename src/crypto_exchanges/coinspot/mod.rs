pub mod structs;
pub mod public_api;
pub mod private_api;

use reqwest::Error;
use serde::de::DeserializeOwned;
use structs::{LatestPrice, LatestPrices, LatestPriceForCoin, TransactionType, TransactionType::{BUY,SELL,SWAP}, Orders, OrderType};

const BASE_URL: &str = "https://www.coinspot.com.au/pubapi/v2/";

async fn get<T>(url: &str) -> Result<T, Error>
where
    T: DeserializeOwned,
{
    let resp = reqwest::get(url).await?.json().await?;
    Ok(resp)
}

pub async fn full_access_status_check() {}
pub async fn my_coin_deposit_address() {}
pub async fn buy_now_quote() {}
pub async fn sell_now_quote() {}
pub async fn place_market_buy_order() {}
pub async fn place_buy_now_order() {}
pub async fn place_market_sell_order() {}
pub async fn place_sell_now_order() {}
pub async fn place_swap_now_order() {}
pub async fn cancel_my_buy_order() {}
pub async fn cancel_my_sell_order() {}
pub async fn get_coin_withdrawal_details() {}
pub async fn coin_withdrawal() {}




pub async fn read_only_status_check() {}
pub async fn open_market_orders() {}
pub async fn completed_market_orders() {}
pub async fn my_coin_balances() {}
pub async fn my_coin_balance() {}
pub async fn my_open_market_orders() {}
pub async fn my_open_limit_orders() {}
pub async fn my_order_history() {}
pub async fn my_market_order_history() {}
pub async fn my_send_and_receive_history() {}
pub async fn my_deposit_history() {}
pub async fn my_withdrawal_history() {}
pub async fn my_affiliate_payments() {}
pub async fn my_referral_payments() {}