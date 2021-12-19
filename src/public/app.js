const vm = new Vue ({
    el: '#vue-instance',
    data () {
        return {
            baseUrl: 'http://localhost:8080/api',
            version : 'v1',
            servers : 0,
            operations : 0,
            keys : 5000,
            x : [],
            yRemap : [],
            yMetric : [],
            value : 0
        }
    },
    methods: {
        async hashring(){
            this.servers = 1;
            this.operations = 1;
            this.x = [],
            this.yRemap = [],
            this.yMetric = []
            await axios.post(`${this.baseUrl}/${this.version}/hashring`,{ keys: this.keys })
        },
        async add(){
            id = (this.servers+1).toString()
            const response = await axios.put(`${this.baseUrl}/${this.version}/server/${id}`)
            this.servers += 1
            this.operations += 1
            this.x.push(this.operations)
            this.yRemap.push(response.data.remap)
            this.yMetric.push(this.keys/this.servers)
            plotRemapChart()
        },
        async remove(){
            id = this.servers.toString()
            const response = await axios.delete(`${this.baseUrl}/${this.version}/server/${id}`)
            this.servers -= 1
            this.operations += 1
            this.x.push(this.operations)
            this.yRemap.push(response.data.remap)
            this.yMetric.push(this.keys/this.servers)
            plotRemapChart()
        },
        async mapping(id){
            const response = await axios.get(`${this.baseUrl}/${this.version}/mapping/${id}`)
            this.value = response
        },
        async mappingAll(){
            const response = await axios.get(`${this.baseUrl}/${this.version}/mapping/all`)
            console.log(response)
        }
    }
})

function plotRemapChart(){
    var x = vm.x;
    var yMetric = vm.yMetric;
    var yRemap = vm.yRemap;
    new Chart("remapChart", {
    type: "line",
    data: {
        labels: x,
        datasets: [
        {
            label : "Remapped Keys",
            data: yRemap,
            borderColor: "red",
            fill: false
        },
        {
            label : "Average Remap Metric",
            data: yMetric,
            borderColor: "green",
            fill: false
        }]
    },
    options: {
        title: {
            display: true,
            text: "Keys Remapped",
            fontSize: 16
        },
        legend: {
            display: true,
            labels: {
                color: 'rgb(255, 99, 132)'
            }}
    }
    });
}

function plotHashRingChart(){

}

