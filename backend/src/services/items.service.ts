import { Request, Response } from "express";
const csv = require("csv-parser");
const fs = require("fs");
const path = require("path");
const jsonexport = require("jsonexport");

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
    let stuff = getStuff(items);
    res.json(stuff.props);
  }

  public async getRandomUniqueProperty(req: Request, res: Response) {
    let items = await getItems("UniqueItems");
    let stuff = getStuff(items);
    let props = stuff.props;
    let randomProp = props[getRandomInt(0, props.length)];
    res.json(randomProp);
  }

  public async randomizeUniques(req: Request, res: Response) {
    let items = await getItems("UniqueItems");
    let stuff = getStuff(items);
    // let randomized = randomizeItems(stuff);
    let csv = await buildCsv(stuff.items.map(x => x.item))

    res.setHeader("Content-Disposition", "attachment; filename=UniqueItems.txt");
    res.send(csv);
  }
}

const buildCsv = async (items : any[]) => {
  let options = {
    rowDelimiter: '\t'
  }
  return new Promise(async (resolve, reject) => {
    await jsonexport(items, options, function(err : any, csv : any){
      if (err) reject(err);
      resolve(csv);
    });
  });

}

const randomizeItems = (stuff: Stuff) => {
  let getRandomProp = () => {
    return stuff.props[getRandomInt(0, stuff.props.length)];
  };
  stuff.items.map((item) => {
    for (var i = 0; i < 12; i++) {
      if (item.item[`prop${i}`] !== "") {
        let randomProp = getRandomProp();
        item.item[`prop${i}`] = randomProp.name;
        item.item[`par${i}`] = randomProp.par;
        item.item[`min${i}`] = randomProp.min;
        item.item[`max${i}`] = randomProp.max;
      }
    }
  });
  return stuff;
};

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

const getStuff = (data: any) => {
  let items: Array<Item> = [];
  let props: Array<Prop> = [];
  data.forEach((item: any) => {
    let propCount = 0;
    let name = item["index"];
    let enabled = item["enabled"];
    if (name !== "") {
      for (var i = 1; i <= 12; i++) {
        let propI = item[`prop${i}`];
        if (propI !== "") {
          propCount++;
          props.push(
            new Prop(propI, item[`par${i}`], item[`min${i}`], item[`max${i}`])
          );
        }
      }
      items.push(new Item(name, propCount, item));
    }
  });
  return new Stuff(props, items);
};

class Stuff {
  props: Array<Prop>;
  items: Array<Item>;

  constructor(props: Array<Prop>, items: Array<Item>) {
    this.props = props;
    this.items = items;
  }
}

class Prop {
  name: string;
  par: string;
  min: string;
  max: string;

  constructor(name: string, par: string, min: string, max: string) {
    this.name = name;
    this.par = par;
    this.min = min;
    this.max = max;
  }
}

class Item {
  name: string;
  numProps: number;
  item: any;

  constructor(name: string, numProps: number, item: any) {
    this.name = name;
    this.numProps = numProps;
    this.item = item;
  }
}

const getRandomInt = (min: number, max: number) => {
  let min_ = Math.ceil(min);
  let max_ = Math.floor(max);
  return Math.floor(Math.random() * (max_ - min_)) + min_;
};
