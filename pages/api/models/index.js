import { openDb } from '../../../lib/db';

export default async function handler(req, res) {
  const db = await openDb();

  if (req.method === 'GET') {
    const models = await db.all('SELECT * FROM models');
    res.status(200).json(models);
  } else {
    res.status(405).json({ message: 'Method not allowed' });
  }
}
