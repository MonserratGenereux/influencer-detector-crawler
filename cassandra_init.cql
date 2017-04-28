# Create Keyspace
CREATE KEYSPACE IF NOT EXISTS influencer_detector
WITH REPLICATION = {'class':'SimpleStrategy', 'replication_factor' : 2}
AND DURABLE_WRITES = true;

# Setup tables
CREATE TABLE IF NOT EXISTS influencer_detector.nodes (
  id BIGINT PRIMARY KEY,
  platform TEXT,
  name TEXT,
  depth bigint,
  description TEXT, 
  categories LIST<TEXT>,
  fan_count BIGINT,
  city TEXT, 
  country TEXT, 
  zip TEXT
);

CREATE TABLE IF NOT EXISTS influencer_detector.edges (
  source BIGINT,
  destination BIGINT,
  PRIMARY KEY (source, destination)
);
