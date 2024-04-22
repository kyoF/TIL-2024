import { RegisterBookApplicatoinService, RegisterBookCommand } from "Application/Book/RegisterBookApplicationService/RegisterBookApplicationService";
import { PrismaBookRepository } from "Infrastructure/Prisma/Book/PrismaBookRepository";
import { PrismaClientManager } from "Infrastructure/Prisma/PrismaClientManager";
import { PrismaTransactionManager } from "Infrastructure/Prisma/PrismaTransactionManager";
import express from "express";

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

    const clientManager = new PrismaClientManager();
    const transactionManager = new PrismaTransactionManager(clientManager);
    const bookRepository = new PrismaBookRepository(clientManager);
    const RegisterBookApplicationService = new RegisterBookApplicatoinService(
      bookRepository,
      transactionManager
    );

    const registerBookCommand: RegisterBookCommand = requestBody;
    await RegisterBookApplicationService.execute(registerBookCommand);

    res.status(200).json({ message: 'success' });
  } catch (error) {
    res.status(500).json({ message: (error as Error).message });
  }
});
