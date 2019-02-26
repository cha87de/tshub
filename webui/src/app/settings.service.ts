import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';

@Injectable({
	providedIn: 'root'
})
export class SettingsService {

	timeframe = '2';

	reloadObservable = new Subject<boolean>();

	constructor() { }


	reload() {
		// trigger reload on any reloadable item
		this.reloadObservable.next(true);
	}
}
