import { openDb } from '../../../lib/db';
import models from '../../../mocks/models.json';
import profiles from '../../../mocks/profiles.json';
import prompts from '../../../mocks/prompts.json';
import scores from '../../../mocks/scores.json';

let isInitialized = false;

export default async function handler(req, res) {
  const db = await openDb();

  if (req.method === 'POST') {
    const { action } = req.body;

    if (action === 'loadMockData') {
      if (!isInitialized) {
        // Initialize the database with mock data
        await db.run('BEGIN TRANSACTION');

        for (const model of models) {
          await db.run(
            'INSERT INTO models (name, path, gpuLayers, ctxSize, batchSize, threads, keep, predict, flashAttn, mlock, cacheTypeK, cacheTypeV) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)',
            [model.name, model.path, model.gpuLayers, model.ctxSize, model.batchSize, model.threads, model.keep, model.predict, model.flashAttn, model.mlock, model.cacheTypeK, model.cacheTypeV]
          );
        }

        for (const profile of profiles) {
          await db.run(
            'INSERT INTO profiles (name, systemPrompt, dryMultiplier, dryBase, dryAllowedLength, dryPenaltyLastN, repeatPenalty, repeatLastN, topK, topP, minP, topA, xtcThreshold, xtcProbability, temperature) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)',
            [profile.name, profile.systemPrompt, profile.dryMultiplier, profile.dryBase, profile.dryAllowedLength, profile.dryPenaltyLastN, profile.repeatPenalty, profile.repeatLastN, profile.topK, profile.topP, profile.minP, profile.topA, profile.xtcThreshold, profile.xtcProbability, profile.temperature]
          );
        }

        for (const prompt of prompts) {
          const profile = profiles.find(p => p.name === prompt.profile);
          await db.run(
            'INSERT INTO prompts (content, solution, profileId, category) VALUES (?, ?, ?, ?)',
            [prompt.content, prompt.solution, profile.id, prompt.category]
          );
        }

        for (const score of scores) {
          await db.run(
            'INSERT INTO scores (modelId, promptId, attempts, score) VALUES (?, ?, ?, ?)',
            [score.modelId, score.promptId, score.attempts, score.score]
          );
        }

        await db.run('COMMIT');
        isInitialized = true;
        res.status(200).json({ message: 'Mock data loaded successfully' });
      } else {
        res.status(200).json({ message: 'Mock data already loaded' });
      }
    } else if (action === 'wipeData') {
      // Wipe all data from the database
      await db.run('BEGIN TRANSACTION');
      await db.run('DELETE FROM scores');
      await db.run('DELETE FROM prompts');
      await db.run('DELETE FROM profiles');
      await db.run('DELETE FROM models');
      await db.run('COMMIT');
      isInitialized = false;
      res.status(200).json({ message: 'Data wiped successfully' });
    } else {
      const { modelId, promptId, attempts } = req.body;
      let score = 0;

      if (attempts === 1) {
        score = 100;
      } else if (attempts === 2) {
        score = 50;
      } else if (attempts === 3) {
        score = 20;
      }

      await db.run(
        'INSERT INTO scores (modelId, promptId, attempts, score) VALUES (?, ?, ?, ?)',
        [modelId, promptId, attempts, score]
      );

      res.status(201).json({ message: 'Score saved successfully' });
    }
  } else if (req.method === 'GET') {
    const scores = await db.all('SELECT * FROM scores');
    res.status(200).json(scores);
  } else {
    res.status(405).json({ message: 'Method not allowed' });
  }
}
