import { Component, OnInit } from '@angular/core';
import { DefaultService, Host } from 'src/app/api';

@Component({
	selector: 'host-list',
	templateUrl: './list.component.html',
	styleUrls: ['./list.component.css']
})
export class ListComponent implements OnInit {

	hosts: Host[];

	constructor(private defaultService: DefaultService) {}
	
	ngOnInit() {
		let hosts = this.defaultService.getHosts();
		hosts.toPromise().then((hosts) => {
			this.hosts = hosts;
		})
	}

}
