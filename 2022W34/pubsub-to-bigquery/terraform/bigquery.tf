###############################################################################
# bigquery: dataset
###############################################################################

resource "google_bigquery_dataset" "daeca" {
  dataset_id = "daeca"
  location   = local.bigquery_location
}
