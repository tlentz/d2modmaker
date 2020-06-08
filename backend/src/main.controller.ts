import { Application } from "express";
import { ItemsService } from "./services/items.service";

export class Controller {
  private itemsService: ItemsService;

  constructor(private app: Application) {
    this.itemsService = new ItemsService();
    this.routes();
  }

  public routes() {
    this.app.route("/api/items/gems").get(this.itemsService.getGems);
    this.app.route("/api/items/runes").get(this.itemsService.getRunes);
    this.app.route("/api/items/setitems").get(this.itemsService.getSetItems);
    this.app.route("/api/items/sets").get(this.itemsService.getSets);
    this.app.route("/api/items/uniqueitems").get(this.itemsService.getUniques);
    this.app
      .route("/api/items/uniqueitems/randomprop")
      .get(this.itemsService.getRandomUniqueProperty);
    this.app
      .route("/api/items/uniqueitems/random")
      .get(this.itemsService.randomizeUniques);
  }
}
