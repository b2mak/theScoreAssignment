import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';

@Injectable()
export class PlayersService {
  testData: string = `[ { "Name": "Ka'Deem Carey", "Team": "CHI", "Position": "RB", "Att_g": 2.7, "Att": 32, "Yds": 126, "Avg": 3.9, "Yds_g": 10.5, "Td": 0, "Lng": "24", "First": 4, "First_percentage": 12.5, "Twenty_plus": 1, "Forty_plus": 0, "Fum": 0 }, { "Name": "Le'Veon Bell", "Team": "PIT", "Position": "RB", "Att_g": 21.8, "Att": 261, "Yds": 1268, "Avg": 4.9, "Yds_g": 105.7, "Td": 7, "Lng": "44", "First": 69, "First_percentage": 26.4, "Twenty_plus": 4, "Forty_plus": 1, "Fum": 3 }, { "Name": "De'Anthony Thomas", "Team": "KC", "Position": "WR", "Att_g": 0.3, "Att": 4, "Yds": 29, "Avg": 7.3, "Yds_g": 2.4, "Td": 0, "Lng": "23", "First": 1, "First_percentage": 25, "Twenty_plus": 1, "Forty_plus": 0, "Fum": 0 } ]`
  playersUrl: string = 'http://localhost:8080/'

  constructor(private http: HttpClient) { }

  getPlayers(
    name: string,
    column: string,
    direction: string,
  ) {
    var queryParams: string[] = []
    if (name !== "") {
      queryParams.push("name=" + name)
    }

    if (column !== "None" && direction !== "None") {
      queryParams.push("orderCol=" + column)
      queryParams.push("orderDirection=" + direction)
    }

    var url = this.playersUrl + "?" + queryParams.join("&")
    return this.http.get<string>(url)
  }

  getPlaceHolder() {
    return JSON.parse(this.testData)
  }

  getTeams() {
    return this.http.get<object[]>('http://localhost:8080/teams')
  }
}