import 'reflect-metadata';
import '../../Program';
import { RegisterBookApplicatoinService, RegisterBookCommand } from "Application/Book/RegisterBookApplicationService/RegisterBookApplicationService";
import express from "express";
import { container } from "tsyringe";

const app = express();
const port = 5000;

app.get('/', (_, res) => {
  res.send('Hello World!');
});

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`);
});

app.use(express.json());
app.post('/book', async (req, res) => {
  try {
    const requestBody = req.body as {
      isbn: string;
      title: string;
      priceAmount: number;
    };

    const registerBookApplicationService = container.resolve(
      RegisterBookApplicatoinService
    );

    const registerBookCommand: RegisterBookCommand = requestBody;
    await registerBookApplicationService.execute(registerBookCommand);

    res.status(200).json({ message: 'success' });
  } catch (error) {
    res.status(500).json({ message: (error as Error).message });
  }
});
