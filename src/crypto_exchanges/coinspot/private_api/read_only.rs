const BASE_URL: &str = "https://www.coinspot.com.au/api/v2/ro/";
use reqwest::Error;

use crate::crypto_exchanges::coinspot::structs::{Status, Orders, MyCoinBalances, MyCoinBalance};

use super::super::get;

pub async fn read_only_status_check() -> Result<Status, Error> {
    get(&format!("{BASE_URL}status")).await
}

pub async fn open_market_orders() -> Result<Orders, Error> {
    get(&format!("{BASE_URL}orders/market/open")).await
}

pub async fn completed_market_orders() -> Result<Orders, Error> {
    get(&format!("{BASE_URL}orders/market/completed")).await
}

pub async fn my_coin_balances() -> Result<MyCoinBalances, Error> {
    get(&format!("{BASE_URL}my/balances")).await
}

pub async fn my_coin_balance(coin_type: String) -> Result<MyCoinBalance, Error> {
    get(&format!("{BASE_URL}my/balances/{coin_type}?available=yes")).await
}

pub async fn my_open_market_orders() {}
pub async fn my_open_limit_orders() {}
pub async fn my_order_history() {}
pub async fn my_market_order_history() {}
pub async fn my_send_and_receive_history() {}
pub async fn my_deposit_history() {}
pub async fn my_withdrawal_history() {}
pub async fn my_affiliate_payments() {}
pub async fn my_referral_payments() {}