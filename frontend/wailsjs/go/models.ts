export namespace system {
	
	export class Asset {
	    assetId: string;
	    assetName: string;
	    itemType: string;
	    manufacturer: string;
	    purchaseDate: string;
	
	    static createFrom(source: any = {}) {
	        return new Asset(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.assetId = source["assetId"];
	        this.assetName = source["assetName"];
	        this.itemType = source["itemType"];
	        this.manufacturer = source["manufacturer"];
	        this.purchaseDate = source["purchaseDate"];
	    }
	}
	export class CPUInfo {
	    manufacturer: string;
	    model: string;
	    cores: number;
	    threads: number;
	    clock: number;
	
	    static createFrom(source: any = {}) {
	        return new CPUInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.manufacturer = source["manufacturer"];
	        this.model = source["model"];
	        this.cores = source["cores"];
	        this.threads = source["threads"];
	        this.clock = source["clock"];
	    }
	}
	export class DiskInfo {
	    manufacturer?: string;
	    model?: string;
	    serial?: string;
	    interface?: string;
	    sizeGB: number;
	    type?: string;
	
	    static createFrom(source: any = {}) {
	        return new DiskInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.manufacturer = source["manufacturer"];
	        this.model = source["model"];
	        this.serial = source["serial"];
	        this.interface = source["interface"];
	        this.sizeGB = source["sizeGB"];
	        this.type = source["type"];
	    }
	}
	export class Disks {
	    drives: DiskInfo[];
	    count: number;
	
	    static createFrom(source: any = {}) {
	        return new Disks(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.drives = this.convertValues(source["drives"], DiskInfo);
	        this.count = source["count"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
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
	export class GPUInfo {
	    manufacturer?: string;
	    model?: string;
	    vramMB: number;
	    driverVersion?: string;
	    clock?: number;
	
	    static createFrom(source: any = {}) {
	        return new GPUInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.manufacturer = source["manufacturer"];
	        this.model = source["model"];
	        this.vramMB = source["vramMB"];
	        this.driverVersion = source["driverVersion"];
	        this.clock = source["clock"];
	    }
	}
	export class GPUs {
	    cards: GPUInfo[];
	    count: number;
	
	    static createFrom(source: any = {}) {
	        return new GPUs(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.cards = this.convertValues(source["cards"], GPUInfo);
	        this.count = source["count"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
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
	export class MotherboardInfo {
	    manufacturer: string;
	    model: string;
	    serial?: string;
	    version?: string;
	
	    static createFrom(source: any = {}) {
	        return new MotherboardInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.manufacturer = source["manufacturer"];
	        this.model = source["model"];
	        this.serial = source["serial"];
	        this.version = source["version"];
	    }
	}
	export class NICInfo {
	    name: string;
	    mac: string;
	    ips?: string[];
	    wifi: boolean;
	
	    static createFrom(source: any = {}) {
	        return new NICInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.mac = source["mac"];
	        this.ips = source["ips"];
	        this.wifi = source["wifi"];
	    }
	}
	export class NetworkInfo {
	    hostname: string;
	    adapters: NICInfo[];
	
	    static createFrom(source: any = {}) {
	        return new NetworkInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.hostname = source["hostname"];
	        this.adapters = this.convertValues(source["adapters"], NICInfo);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
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
	export class RAMModule {
	    manufacturer?: string;
	    model?: string;
	    capacityMB: number;
	    speedMHz?: number;
	    ddr?: string;
	
	    static createFrom(source: any = {}) {
	        return new RAMModule(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.manufacturer = source["manufacturer"];
	        this.model = source["model"];
	        this.capacityMB = source["capacityMB"];
	        this.speedMHz = source["speedMHz"];
	        this.ddr = source["ddr"];
	    }
	}
	export class RAMInfo {
	    modules: RAMModule[];
	    totalMB: number;
	    count: number;
	
	    static createFrom(source: any = {}) {
	        return new RAMInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.modules = this.convertValues(source["modules"], RAMModule);
	        this.totalMB = source["totalMB"];
	        this.count = source["count"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
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

}

