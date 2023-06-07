mod crypto_exchanges;

#[tokio::main]
async fn main() {
    match crypto_exchanges::coinspot::get_latest_prices(None, None).await {
        Ok(prices) => {
            // Handle successful response
            println!("{:#?}", prices);
        }
        Err(err) => {
            // Handle error
            eprintln!("Error fetching: {}", err);
        }
    }

    match crypto_exchanges::coinspot::get_latest_prices(Some(String::from("BTC")), None).await {
        Ok(prices) => {
            // Handle successful response
            println!("{:#?}", prices);
        }
        Err(err) => {
            // Handle error
            eprintln!("Error fetching: {}", err);
        }
    }
}
