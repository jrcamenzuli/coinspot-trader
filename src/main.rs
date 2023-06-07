mod crypto_exchanges;

#[tokio::main]
async fn main() {
    match crypto_exchanges::coinspot::get_latest_prices(None,None).await {
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
