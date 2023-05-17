(ns coinspot-trader
  (:require
    [clj-http.client :as client]
    [clj-time.coerce :as coerce]))

(defn get-snapshots [t]
  (let [url "http://192.168.0.40:10000/query"
        response (client/get url {:as :json, :query-params {"t" t}})
        snapshots (mapv #(update % :Time coerce/from-string) (:body response))]
    snapshots))

(get-snapshots "1684323847")
