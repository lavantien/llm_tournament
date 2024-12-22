import { open } from 'sqlite';
import sqlite3 from 'sqlite3';

export async function openDb() {
  const db = await open({
    filename: './database.sqlite',
    driver: sqlite3.Database
  });

  // Create models table if it doesn't exist
  await db.exec(`
    CREATE TABLE IF NOT EXISTS models (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      name TEXT NOT NULL,
      path TEXT NOT NULL,
      gpuLayers INTEGER,
      ctxSize INTEGER,
      batchSize INTEGER,
      threads INTEGER,
      keep INTEGER,
      predict INTEGER,
      flashAttn BOOLEAN,
      mlock BOOLEAN,
      cacheTypeK TEXT,
      cacheTypeV TEXT
    );
  `);

  // Create profiles table if it doesn't exist
  await db.exec(`
    CREATE TABLE IF NOT EXISTS profiles (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      name TEXT NOT NULL,
      systemPrompt TEXT,
      dryMultiplier REAL,
      dryBase REAL,
      dryAllowedLength INTEGER,
      dryPenaltyLastN INTEGER,
      repeatPenalty REAL,
      repeatLastN INTEGER,
      topK INTEGER,
      topP REAL,
      minP REAL,
      topA REAL,
      xtcThreshold REAL,
      xtcProbability REAL,
      temperature REAL
    );
  `);

  // Create prompts table if it doesn't exist
  await db.exec(`
    CREATE TABLE IF NOT EXISTS prompts (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      content TEXT NOT NULL,
      solution TEXT,
      profileId INTEGER,
      category TEXT,
      FOREIGN KEY (profileId) REFERENCES profiles(id)
    );
  `);

  // Create scores table if it doesn't exist
  await db.exec(`
    CREATE TABLE IF NOT EXISTS scores (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      modelId INTEGER,
      promptId INTEGER,
      attempts INTEGER DEFAULT 0,
      score REAL,
      FOREIGN KEY (modelId) REFERENCES models(id),
      FOREIGN KEY (promptId) REFERENCES prompts(id)
    );
  `);

  return db;
}
