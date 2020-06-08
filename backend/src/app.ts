import express, { Application } from "express";
import { MONGO_URL } from "./constants";
import bodyParser from "body-parser";
import cors from "cors";
import mongoose from "mongoose";
import { Controller } from "./main.controller";

class App {
  public app: Application;
  public itemsController: Controller;

  constructor() {
    this.app = express();
    this.setConfig();
    // this.setMongoConfig();
    this.itemsController = new Controller(this.app);
  }

  private setConfig() {
    this.app.use(bodyParser.json({ limit: "50mb" }));
    this.app.use(bodyParser.urlencoded({ limit: "50mb", extended: true }));
    this.app.use(cors());
    this.app.set("json spaces", 4);
  }

  //Connecting to our MongoDB database
  private setMongoConfig() {
    mongoose.Promise = global.Promise;
    mongoose.connect(MONGO_URL, {
      useNewUrlParser: true,
    });
  }
}

export default new App().app;
