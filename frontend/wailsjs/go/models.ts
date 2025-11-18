export namespace main {
	export class TridFileTypeResult {
		name: string;
		extension: string;
		mimeType: string;
		description: string;
		confidence: number;
		url: string;
		remarks: string;

		static createFrom(source: any = {}) {
			return new TridFileTypeResult(source);
		}

		constructor(source: any = {}) {
			if ('string' === typeof source) source = JSON.parse(source);
			this.name = source['name'];
			this.extension = source['extension'];
			this.mimeType = source['mimeType'];
			this.description = source['description'];
			this.confidence = source['confidence'];
			this.url = source['url'];
			this.remarks = source['remarks'];
		}
	}
	export class TridScanResult {
		fileName: string;
		fileSize: number;
		matches: TridFileTypeResult[];
		bestMatch?: TridFileTypeResult;
		totalMatches: number;
		error?: string;

		static createFrom(source: any = {}) {
			return new TridScanResult(source);
		}

		constructor(source: any = {}) {
			if ('string' === typeof source) source = JSON.parse(source);
			this.fileName = source['fileName'];
			this.fileSize = source['fileSize'];
			this.matches = this.convertValues(source['matches'], TridFileTypeResult);
			this.bestMatch = this.convertValues(source['bestMatch'], TridFileTypeResult);
			this.totalMatches = source['totalMatches'];
			this.error = source['error'];
		}

		convertValues(a: any, classs: any, asMap: boolean = false): any {
			if (!a) {
				return a;
			}
			if (a.slice && a.map) {
				return (a as any[]).map((elem) => this.convertValues(elem, classs));
			} else if ('object' === typeof a) {
				if (asMap) {
					for (const key of Object.keys(a)) {
						a[key] = new classs(a[key]);
					}
					return a;
				}
				return new classs(a);
			}
			return a;
		}
	}
	export class TridUpdateInfo {
		currentMD5: string;
		latestMD5: string;
		isUpToDate: boolean;
		lastUpdated: string;
		defsCount: number;
		error?: string;

		static createFrom(source: any = {}) {
			return new TridUpdateInfo(source);
		}

		constructor(source: any = {}) {
			if ('string' === typeof source) source = JSON.parse(source);
			this.currentMD5 = source['currentMD5'];
			this.latestMD5 = source['latestMD5'];
			this.isUpToDate = source['isUpToDate'];
			this.lastUpdated = source['lastUpdated'];
			this.defsCount = source['defsCount'];
			this.error = source['error'];
		}
	}
	export class UpdateInfo {
		currentVersion: string;
		latestVersion: string;
		updateAvailable: boolean;
		releaseUrl: string;
		releaseNotes: string;
		publishedAt: string;

		static createFrom(source: any = {}) {
			return new UpdateInfo(source);
		}

		constructor(source: any = {}) {
			if ('string' === typeof source) source = JSON.parse(source);
			this.currentVersion = source['currentVersion'];
			this.latestVersion = source['latestVersion'];
			this.updateAvailable = source['updateAvailable'];
			this.releaseUrl = source['releaseUrl'];
			this.releaseNotes = source['releaseNotes'];
			this.publishedAt = source['publishedAt'];
		}
	}
}
