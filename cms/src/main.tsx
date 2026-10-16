import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './interfaces/App.tsx'
import './index.css'
import type { ApplicationServices } from "@/application/services";
import { Service as AuthService } from "@/application/service/auth/service";
import { Service as MessageService } from "@/application/service/message/service";
import { Service as NotificationService } from "@/application/service/notification/service";
import { Service as TransactionService } from "@/application/service/transaction/service";
import { FirebaseRepository as AuthRepository } from "@/infrastructure/auth/repository";
import { HttpRepository as MessageRepository } from "@/infrastructure/http/message/repository";
import { HttpRepository as NotificationRepository } from "@/infrastructure/http/notification/repository";
import { httpClient } from "@/infrastructure/http/client";
import { HttpRepository as TransactionRepository } from "@/infrastructure/http/transaction/repository";

const application: ApplicationServices = {
  auth: new AuthService(new AuthRepository()),
  message: new MessageService(new MessageRepository(httpClient)),
  notification: new NotificationService(new NotificationRepository(httpClient)),
  transaction: new TransactionService(new TransactionRepository(httpClient)),
};

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <App application={application} />
  </React.StrictMode>,
)
