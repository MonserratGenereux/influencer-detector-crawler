# Create Keyspace
CREATE KEYSPACE IF NOT EXISTS influencer_detector
WITH REPLICATION = {'class':'SimpleStrategy', 'replication_factor' : 2}
AND DURABLE_WRITES = true;

# Setup tables

CREATE TABLE IF NOT EXISTS influencer_detector.nodes (
  id TEXT PRIMARY KEY,
  platform TEXT,
  name TEXT,
  description TEXT, 
  categories LIST<TEXT>,
  fan_count BIGINT,
  city TEXT, 
  country TEXT, 
  zip TEXT
);

CREATE TABLE IF NOT EXISTS influencer_detector.edges (
  source TEXT,
  destination TEXT,
  PRIMARY KEY (source, destination)
);

# Insert references
INSERT INTO influencer_detector.nodes
  ( id, platform, name, description, categories, fan_count, city, country, zip)
VALUES ( 
  '683165801724841', 
  'facebook', 
  'Wizeline Mexico', 
  'Wizeline transforms how business design, develop and deliver technology products. Join our team! Now hiring in GDL.', 
  ['Internet Company', 'Local Service', 'Professional Service'],
  3684,
  'Zapopan',
  'Mexico',
  '45050'
  );

INSERT INTO influencer_detector.edges ( source, destination )
VALUES ( '683165801724841', '683165801724841');