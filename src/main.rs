mod crypto_exchanges;

use crypto_exchanges::coinspot::structs::TransactionType::{BUY, SELL};

#[tokio::main]
async fn main() {
    match crypto_exchanges::coinspot::get_latest_prices().await {
        Ok(prices) => {
            // Handle successful response
            println!("{:#?}", prices);
        }
        Err(err) => {
            // Handle error
            eprintln!("Error fetching: {}", err);
        }
    };

    match crypto_exchanges::coinspot::get_latest_price(String::from("BTC"), None).await {
        Ok(prices) => {
            // Handle successful response
            println!("{:#?}", prices);
        }
        Err(err) => {
            // Handle error
            eprintln!("Error fetching: {}", err);
        }
    };

    match crypto_exchanges::coinspot::get_latest_price(
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

    match crypto_exchanges::coinspot::get_latest_transaction_price(
        String::from("BTC"),
        None,
        Some(BUY),
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

    match crypto_exchanges::coinspot::get_latest_transaction_price(
        String::from("BTC"),
        None,
        Some(SELL),
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
