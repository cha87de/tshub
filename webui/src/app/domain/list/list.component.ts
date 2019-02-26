import { Component, OnInit } from '@angular/core';
import { DefaultService, Domain } from 'src/app/api';
import { ActivatedRoute } from '@angular/router';
import 'rxjs/add/operator/filter';
import { SettingsService } from 'src/app/settings.service';
import { Subscription } from 'rxjs';

@Component({
	selector: 'domain-list',
	templateUrl: './list.component.html',
	styleUrls: ['./list.component.css']
})
export class ListComponent implements OnInit {

	domains: Domain[];
	hostname: string = "";

	reloadSubscription: Subscription;

	constructor(private route: ActivatedRoute, private defaultService: DefaultService, private settingsService: SettingsService) {}
	
	ngOnInit() {
		this.route.queryParams
		.filter(params => params.hostname)
		.subscribe(params => {
			this.hostname = params.hostname;
			this.queryDomains();
		});

		if (this.hostname == ''){
			this.queryDomains();
		}

		// register to reload event
		this.reloadSubscription = this.settingsService.reloadObservable.subscribe(
			reload => {
				this.queryDomains();
			}
		);		
	}

	ngOnDestroy(){
		this.reloadSubscription.unsubscribe();
	}

	queryDomains() {
		let domains = this.defaultService.getDomains(this.hostname);
		domains.toPromise().then((domains) => {
			this.domains = domains;
		})
	}

}
