<template>
    <div>
      <label for="options">Choose an option:</label>
      <select v-model="selectedOption" id="options">
        <option disabled value="">Please select one</option>
        <option v-for="(option, index) in options" :key="index" :value="option">
          {{ option }}
        </option>
      </select>
  
      <p>Selected option: {{ selectedOption }}</p>
    </div>
  </template>
  
  <script>
  import { ref } from 'vue';
  
  
  export default {
    setup() {
      const selectedOption = ref(""); // The default value for the select box
      const options = ref([
        "Option 1",
        "Option 2",
        "Option 3",
        "Option 4",
      ]);
  
      return {
        selectedOption,
        options,
      };
    },
  };


  export default {
  name: 'DataView',
  data() {
    return {
      data: [],
      searchTerm: '',
      //searchResults: []
    };
  },
  mounted() {
    this.search();
  },
  methods: {

    search() {

      axios.get('https://data.gov.sg/api/action/datastore_search?resource_id=eb8b932c-503c-41e7-b513-114cffbe2338&limit=20', {
        params: {

          q: this.searchTerm
        }
      })
        .then(response => {
          this.data =  response.data.result.records;
        })
        .catch(error => {
          console.error(error);
        });
    }

  },
};
  </script>
  

  