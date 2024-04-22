import { ITransactionManager } from "Application/shared/ITransactionManager";
import { PrismaClientManager } from "./PrismaClientManager";
import prisma from "./PrismaClient";
import { inject, injectable } from "tsyringe";

@injectable()
export class PrismaTransactionManager implements ITransactionManager {
  constructor(
    @inject('IDataAccessClientManager')
    private clienManager: PrismaClientManager
  ) { }

  async begin<T>(callback: () => Promise<T>): Promise<T | undefined> {
    return await prisma.$transaction(async (transaction) => {
      this.clienManager.setClient(transaction)

      const res = await callback();

      this.clienManager.setClient(prisma);

      return res;
    });
  }
}
