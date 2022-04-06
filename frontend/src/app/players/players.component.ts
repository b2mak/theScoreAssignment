import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import { PlayersService } from './players.service';
import { AngularCsv } from 'angular-csv-ext/dist/Angular-csv';

@Component({
  selector: 'app-players',
  templateUrl: './players.component.html',
  styleUrls: ['./players.component.css']
})
export class PlayersComponent implements OnInit {
  name: string
  column: string
  direction: string
  tableColumns: string[] = [
    "Name",
    "Team",
    "Position",
    "Att/g",
    "Att",
    "Yds",
    "Avg",
    "Yds/g",
    "TD",
    "Lng",
    "First",
    "First%",
    "Twenty+",
    "Forty+",
    "Fumbles",
  ];
  players

  constructor(private service: PlayersService) {
    this.players = service.getPlaceHolder()
    this.name = ""
    this.column = "None"
    this.direction = "None"
  }

  ngOnInit(): void {
    this.showPlayers();
  }

  showPlayers() {
    this.service.getPlayers(
      this.name,
      this.column,
      this.direction,
    ).subscribe(
      data => {
        this.players = data
      }
    )
  }

  updateTable() {
    this.showPlayers();
  }

  download() {
    var columnTitles: {[k: string]: string} = {}
    this.tableColumns.forEach(element => {
      columnTitles[element] = element
    });
    var data = [columnTitles].concat(this.players)
    new AngularCsv(data, 'download');
  }
}
