import { openDb } from '../../../lib/db';

export default async function handler(req, res) {
  const db = await openDb();

  if (req.method === 'GET') {
    const profiles = await db.all('SELECT * FROM profiles');
    res.status(200).json(profiles);
  } else {
    res.status(405).json({ message: 'Method not allowed' });
  }
}
