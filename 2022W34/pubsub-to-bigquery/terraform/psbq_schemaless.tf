
###############################################################################
# pubsub topic
###############################################################################

resource "google_pubsub_topic" "daeca_schemaless" {
  name                       = "daeca-schemaless"
  message_retention_duration = "${7 * 24 * 60 * 60}s"
}

resource "google_pubsub_subscription" "daeca_schemaless_pull" {
  name                       = "daeca-schemaless-pull"
  topic                      = google_pubsub_topic.daeca_schemaless.id
  ack_deadline_seconds       = 600
  message_retention_duration = "${7 * 24 * 60 * 60}s"
  retry_policy {
    minimum_backoff = "1s"
    maximum_backoff = "60s"
  }
  dead_letter_policy {
    dead_letter_topic     = google_pubsub_topic.daeca_schemaless_pull_dl.id
    max_delivery_attempts = 10
  }
}

resource "google_pubsub_topic" "daeca_schemaless_pull_dl" {
  name                       = "daeca-schemaless-pull-dl"
  message_retention_duration = "${7 * 24 * 60 * 60}s"
}

resource "google_pubsub_subscription" "daeca_schemaless_pull_dl" {
  name                       = "daeca-schemaless-pull-dl"
  topic                      = google_pubsub_topic.daeca_schemaless_pull_dl.id
  ack_deadline_seconds       = 600
  message_retention_duration = "${7 * 24 * 60 * 60}s"
}

###############################################################################
# bigquery: table
###############################################################################

resource "google_bigquery_table" "daeca_schemaless" {
  dataset_id = google_bigquery_dataset.daeca.dataset_id
  table_id   = "daeca-schemaless"

  time_partitioning {
    field = "publish_time"
    type  = "DAY"
  }

  deletion_protection = false

  schema = file("../spec/bigquery/pubsub_bigquery_schemaless.json")
}

###############################################################################
# bigquery subscription
###############################################################################

resource "google_pubsub_subscription" "daeca_schemaless_bq" {
  name                       = "daeca-schemaless-bq"
  topic                      = google_pubsub_topic.daeca_schemaless.id
  ack_deadline_seconds       = 600
  message_retention_duration = "${7 * 24 * 60 * 60}s"
  bigquery_config {
    table          = "${google_bigquery_table.daeca_schemaless.project}.${google_bigquery_table.daeca_schemaless.dataset_id}.${google_bigquery_table.daeca_schemaless.table_id}"
    write_metadata = true
  }
  retry_policy {
    minimum_backoff = "1s"
    maximum_backoff = "60s"
  }
  dead_letter_policy {
    dead_letter_topic     = google_pubsub_topic.daeca_schemaless_pull_dl.id
    max_delivery_attempts = 10
  }
}

resource "google_pubsub_topic" "daeca_schemaless_bq_dl" {
  name                       = "daeca-schemaless-bq-dl"
  message_retention_duration = "${7 * 24 * 60 * 60}s"
}

resource "google_pubsub_subscription" "daeca_schemaless_bq_dl" {
  name                       = "daeca-schemaless-bq-dl"
  topic                      = google_pubsub_topic.daeca_schemaless_bq_dl.id
  ack_deadline_seconds       = 600
  message_retention_duration = "${7 * 24 * 60 * 60}s"
}
