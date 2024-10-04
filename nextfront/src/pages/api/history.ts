import type { NextApiRequest, NextApiResponse } from 'next';
import fs from 'fs';
import path from 'path';

export default function handler(req: NextApiRequest, res: NextApiResponse) {
  const filePath = path.join(process.cwd(), 'data', 'videoHistory.json');

  const dirPath = path.dirname(filePath);
  if (!fs.existsSync(dirPath)) {
    fs.mkdirSync(dirPath, { recursive: true });
  }

  if (req.method === 'GET') {
    if (fs.existsSync(filePath)) {
      const fileContent = fs.readFileSync(filePath, 'utf8');
      const data = fileContent ? JSON.parse(fileContent) : [];
      res.status(200).json(data);
    } else {
      res.status(404).json({ message: 'Historial no encontrado' });
    }
  } else if (req.method === 'POST') {
    const data = req.body;

    let fileData = [];
    if (fs.existsSync(filePath)) {
      const fileContent = fs.readFileSync(filePath, 'utf8');
      fileData = fileContent ? JSON.parse(fileContent) : [];
    }

    fileData.push(data);

    fs.writeFileSync(filePath, JSON.stringify(fileData, null, 2));

    res.status(200).json({ message: 'Historial guardado exitosamente' });
  } else {
    res.status(405).json({ message: 'Método no permitido' });
  }
}
