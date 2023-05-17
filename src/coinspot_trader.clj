(ns coinspot-trader
  (:require [clojure.java.io :as io]
            [clojure.string :as str]))

(defn make-get-request [url params]
  (let [query-string (str/join "&" (for [[k v] params] (str (subs (str k) 1) "=" v)))]
    (println "GET:" (str url "?" query-string))
    (with-open [reader (io/reader (java.net.URL. (str url "?" query-string)))]
      (str/join (line-seq reader)))))


(let [url "http://192.168.0.40:10000/query"
      params {:t "1684162058"
              :farts "farting"}]
  (while true
    (println (make-get-request url params))
    (Thread/sleep 5000)))