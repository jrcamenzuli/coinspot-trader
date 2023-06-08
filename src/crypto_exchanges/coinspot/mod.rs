pub mod private_api;
pub mod public_api;
pub mod structs;

use reqwest::Error;
use serde::de::DeserializeOwned;

async fn get<T>(url: &str) -> Result<T, Error>
where
    T: DeserializeOwned,
{
    let resp = reqwest::get(url).await?.json().await?;
    Ok(resp)
}
