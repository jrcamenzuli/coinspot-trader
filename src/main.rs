mod crypto_exchanges;

use crypto_exchanges::coinspot::structs::TransactionType::{BUY, SELL};
use crypto_exchanges::coinspot::structs::OrderType::{OPEN, COMPLETED};
use crypto_exchanges::coinspot::public_api;

#[tokio::main]
async fn main() {
    match public_api::get_latest_prices().await {
        Ok(prices) => {
            // Handle successful response
            println!("{:#?}", prices);
        }
        Err(err) => {
            // Handle error
            eprintln!("Error fetching: {}", err);
        }
    };

    match public_api::get_latest_price(String::from("BTC"), None).await {
        Ok(prices) => {
            // Handle successful response
            println!("{:#?}", prices);
        }
        Err(err) => {
            // Handle error
            eprintln!("Error fetching: {}", err);
        }
    };

    match public_api::get_latest_price(
        String::from("BTC"),
        Some(String::from("USDT")),
    )
    .await
    {
        Ok(prices) => {
            // Handle successful response
            println!("{:#?}", prices);
        }
        Err(err) => {
            // Handle error
            eprintln!("Error fetching: {}", err);
        }
    };

    match public_api::get_latest_buy_price(String::from("BTC"))
    .await
    {
        Ok(prices) => {
            // Handle successful response
            println!("{:#?}", prices);
        }
        Err(err) => {
            // Handle error
            eprintln!("Error fetching: {}", err);
        }
    };

    match public_api::get_latest_sell_price(String::from("BTC"))
    .await
    {
        Ok(prices) => {
            // Handle successful response
            println!("{:#?}", prices);
        }
        Err(err) => {
            // Handle error
            eprintln!("Error fetching: {}", err);
        }
    };

    match public_api::get_orders(
        String::from("BTC"),
        OPEN,
        None
    )
    .await
    {
        Ok(prices) => {
            // Handle successful response
            println!("{:#?}", prices);
        }
        Err(err) => {
            // Handle error
            eprintln!("Error fetching: {}", err);
        }
    };
}
