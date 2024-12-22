import path from 'path';
import fs from 'fs';

export default function handler(req, res) {
  const { filename } = req.query;
  const filePath = path.join(process.cwd(), 'mocks', `${filename}.json`);

  fs.readFile(filePath, 'utf8', (err, data) => {
    if (err) {
      return res.status(500).json({ error: 'Error reading file' });
    }
    res.setHeader('Content-Type', 'application/json');
    res.status(200).json(JSON.parse(data));
  });
}
