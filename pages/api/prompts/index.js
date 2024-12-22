import { openDb } from '../../../lib/db';

export default async function handler(req, res) {
  const db = await openDb();

  if (req.method === 'GET') {
    const prompts = await db.all('SELECT * FROM prompts');
    res.status(200).json(prompts);
  } else {
    res.status(405).json({ message: 'Method not allowed' });
  }
}
