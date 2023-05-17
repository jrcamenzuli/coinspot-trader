(ns coinspot-trader
  (:require [clojure.java.io :as io]
            [clojure.string :as str]
            [clj-http.client :as client]
            [cheshire.core :as json]))
(defn get-snapshot [t]
  (let
    [
     url "http://192.168.0.40:10000/query"
     response (client/get url {:as :reader :accept :json :query-params {"t" t}})
     ]
    (with-open [reader (:body response)]  ; closes the underlying connection when we're done
      (let [snapshots (json/parse-stream reader true)]
        ; You must perform all reads from the stream inside `with-open`,
        ; any , any lazy
        (doall (for [snapshot snapshots]
                 snapshot))))))
(get-snapshot "1684323847")