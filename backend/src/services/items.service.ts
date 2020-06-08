import { Request, Response } from "express";
import { resolve } from "dns";
import { rejects } from "assert";
const csv = require("csv-parser");
const fs = require("fs");
const path = require("path");

export class ItemsService {
  constructor() {}
  public async getGems(req: Request, res: Response) {
    let items = await getItems("Gems");
    res.json(items);
  }

  public async getRunes(req: Request, res: Response) {
    let items = await getItems("Runes");
    res.json(items);
  }

  public async getSets(req: Request, res: Response) {
    let items = await getItems("Sets");
    res.json(items);
  }

  public async getSetItems(req: Request, res: Response) {
    let items = await getItems("SetItems");
    res.json(items);
  }

  public async getUniques(req: Request, res: Response) {
    let items = await getItems("UniqueItems");
    res.json(items);
  }
}

const getItems = async (fname: string) => {
  return new Promise(async (resolve, reject) => {
    let results: any = [];
    let fpath = path.join(__dirname, `../assets/d2-txts/${fname}.txt`);
    await fs
      .createReadStream(fpath)
      .pipe(csv({ separator: "\t" }))
      .on("data", (data: any) => results.push(data))
      .on("end", () => {
        console.log(`Done fetching ${fname}.txt`);
        resolve(results);
      });
  });
};
