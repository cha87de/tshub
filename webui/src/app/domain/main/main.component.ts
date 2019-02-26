import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'domain-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.css']
})
export class MainComponent implements OnInit {

  hostname: string;

  constructor(private route: ActivatedRoute) { }
	
	ngOnInit() {
		this.route.queryParams
		.filter(params => params.hostname)
		.subscribe(params => {
			this.hostname = params.hostname;
		});
	}

}
