const vm = new Vue ({
    el: '#vue-instance',
    data () {
        return {
            baseUrl: 'http://localhost:8080/api',
            version : 'v1',
            operations : 0,
            value : "NIL"
        }
    },
    methods: {
        async add(id){
            const response = await axios.post(`${this.baseUrl}/${this.version}/add`, { id: id })
            value = response.id
        },
        async remove(id){
            const response = await axios.delete(`${this.baseUrl}/${this.version}/remove`, { id: id })
            value = response.id
        }
    }
})
