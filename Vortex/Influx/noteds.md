[
  {
    "operation": "shift",
    "spec": {
      // match "rating-primary" explicitly
      //"time": 1559630455,
      // "tags": { "test_tag": "test" },
      // match the rest using the * wildcard.
      "*": "fields.&"
        // Assuming "rating-Price" :
        // "SecondaryRatings.&" = "rating-Price"   // the whole key, syntatic sugar version
        // "SecondaryRatings.&0" = "rating-Price"  // the whole key, explict about how far up the tree to look, aka 0
        // "SecondaryRatings.&(0,1)" = "Price"     // the first * in the key &0 levels up the intput tree
    }
  },
  {
    "operation": "default",
    "spec": {
      "id": "Y.F.${ticker}",
      "measurement": "quarter-fiancial",
      "time": 1559630455,
      "tags": { "topic_tag": "${ticker}" }
    }
}