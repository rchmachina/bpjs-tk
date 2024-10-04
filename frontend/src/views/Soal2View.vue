<template>
  <div class="p-20 ml-20 overflow-x-auto">
    <label for="itemSelect">Select the parent:</label>
    <div class="flex items-center">
      <select id="itemSelect" v-model="selectedId" @change="handleSelectChange">
        <option disabled value="">Please select one</option> <!-- Optional placeholder -->
        <option v-for="item in data" :key="item.id" :value="item.id">
          {{ item.nameData }}
        </option>
      </select>
      <button @click="removeSelected" class="ml-2 bg-red-800 hover:text-red-800">
        Clear
      </button>
    </div>

    
      <table class="table w-full">
        <thead>
          <tr>
            <th class="border px-4 py-2">id</th>
            <th class="border px-4 py-2">nama buah</th>

          </tr>
        </thead>
        <tbody>
          <tr v-for="item in data" :key="item.id">
            <td class="border px-4 py-2">{{ item.id }}</td>
            <td class="border px-4 py-2">{{ item.nameData }}</td>
          </tr>
        </tbody>
      </table>

  </div>
</template>

<script>
import { ref, onMounted } from "vue";
import axios from "axios";

export default {
  name: "DataView",
  setup() {
    const data = ref([]);
    const selectedId = ref("");
    const searchTerm = ref("");


    const handleSelectChange = () => {
      if (selectedId.value !== '') {
        searchData(); // Call searchData when selectedId changes
      } else {
        fetchingGetAllData(); // Call to fetch all data if no selection
      }
    };
    // Method to search for data based on the search term


    const removeSelected = () => {
      console.log(`Removing item with ID: ${selectedId.value}`);
      // Logic to remove the selected item can go here
      // e.g., send a DELETE request to your API if needed

      // Optionally reset selectedId
      selectedId.value = "";
      fetchingGetAllData(); // Refresh data if necessary
    };

    const searchData = () => {
      const link = `${import.meta.env.VITE_BACKEND_API_URL}/get-child-soal2?id=${selectedId.value}`;
      console.log(link);

      axios
        .get(link)
        .then((response) => {
          data.value = response.data.data;
          console.log(data.value);
        })
        .catch((error) => {
          console.error(error);
        });
    };

    // Method to fetch all data
    const fetchingGetAllData = () => {
      const link = `${import.meta.env.VITE_BACKEND_API_URL}/get-data-soal2`;
      console.log(link);

      axios
        .get(link)
        .then((response) => {
          data.value = response.data.data;
          console.log(data.value);
        })
        .catch((error) => {
          console.error(error);
        });
    };

    // Lifecycle hook to fetch all data when the component is mounted
    onMounted(() => {
      fetchingGetAllData();
    });

    return {
      data,
      selectedId,
      handleSelectChange,
      fetchingGetAllData,
      removeSelected
    };
  },
};
</script>

<style scoped>
/* No additional styles needed for DaisyUI table */
</style>
