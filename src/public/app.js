const vm = new Vue ({
    el: '#vue-instance',
    data () {
        return {
            baseUrl: 'http://localhost:8080/api',
            version : 'v1',
            nodes : 0,
            operations : 0,
            keys : 50,
            MAX : 4294967295,
            x : [],
            y : [],
            value : 0
        }
    },
    methods: {
        async hashring(){
            this.nodes = 0;
            this.operations = 0;
            this.value = this.keys/this.MAX
            await axios.post(`${this.baseUrl}/${this.version}/hashring`,{ keys: keys })
        },
        async add(){
            this.nodes += 1
            id = this.nodes.toString()
            const response = await axios.put(`${this.baseUrl}/${this.version}/node/${id}`)
            this.operations += 1
            this.x.push(this.operations)
            this.y.push(response.data.remap)
        },
        async remove(){
            id = this.nodes.toString()
            const response = await axios.delete(`${this.baseUrl}/${this.version}/node/${id}`)
            this.operations += 1
            this.x.push(this.operations)
            this.y.push(response.data.remap)
            this.nodes -= 1
        },
        async mapping(id){
            const response = await axios.get(`${this.baseUrl}/${this.version}/mapping/${id}`)
            this.value = response
        },
        async mappingAll(){
            const response = await axios.get(`${this.baseUrl}/${this.version}/mapping/all`)
            this.value = response
        }
    }
})
