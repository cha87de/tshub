import { Component, OnInit } from '@angular/core';
import { SettingsService } from '../settings.service';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit {

  timeframe: string;

  constructor(private settingsService: SettingsService) {  }

  ngOnInit() {
    this.timeframe = this.settingsService.timeframe;
  }

  onTimeframeChange(event) {
    this.settingsService.timeframe = this.timeframe;
    this.reload();    
  }

  reload(){
    this.settingsService.reload();
  }

}
