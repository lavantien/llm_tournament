import { openDb } from '../../../lib/db';

export default async function handler(req, res) {
  const db = await openDb();

  if (req.method === 'POST') {
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
  } else if (req.method === 'GET') {
    const scores = await db.all('SELECT * FROM scores');
    res.status(200).json(scores);
  } else {
    res.status(405).json({ message: 'Method not allowed' });
  }
}
