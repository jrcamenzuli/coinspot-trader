(ns coinspot-trader
  (:require
    [clj-http.client :as client]
    [clj-time.coerce :as coerce]
    [clojure.core.async :as async :refer [<! go-loop]]))

(defn get-snapshots [t]
  (let [url "http://192.168.0.40:10000/query"
        response (client/get url {:as :json, :query-params {"t" t}})
        snapshots (mapv #(update % :Time coerce/from-string) (:body response))]
    snapshots))

(defn process-snapshots [snapshots]
  (println snapshots)
  )

(defn fetch-snapshots-loop []
  "The go-loop construct creates an infinite loop with recur that can be paused and resumed asynchronously using
  channels. In this case, after fetching and processing snapshots, the code waits for 5 seconds using the async/timeout
  function and then resumes the loop by calling recur."
  (go-loop []
           (let [snapshots (get-snapshots "1684330942")]
             ;; Do something with the snapshots, e.g., process or print them
             (process-snapshots snapshots))
           (async/<! (async/timeout 5000))  ; Wait for 5 seconds
           (recur)))
(fetch-snapshots-loop)
