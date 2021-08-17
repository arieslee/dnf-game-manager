<template>
    <div class="role-index">
        <v-container fill-height fluid class="align-start">
            <v-row class="mx-auto pt-6" width="90%">
                <v-col>
                    <div class="page-title">
                        角色列表
                    </div>
                    <v-data-table
                        :headers="headers"
                        :items="dataList"
                        :items-per-page="pageSize"
                        :loading="loading"
                        :hide-default-footer="true"
                        :disable-pagination="true"
                        loading-text="数据载入中..."
                        no-data-text="暂无数据"
                        class="elevation-1"
                    ></v-data-table>
                </v-col>
            </v-row>
        </v-container>
    </div>
</template>

<script>
import {getRoleList} from "@/api/account";

export default {
    data() {
        return {
            loading:false,
            headers: [
                {
                    text: '角色ID',
                    align: 'start',
                    sortable: false,
                    value: 'charac_no',
                },
                {text: '角色名', value: 'charac_name'},
                {text: '等级', value: 'lev'},
            ],
            dataList:[],
            pageSize: Number(process.env.VUE_APP_PAGE_SIZE),
            totalCount:0,
            currentPage:1,
            totalPage:1,
        }
    },
    computed:{
        uid() {
            return this.$store.getters.uid;
        },
    },
    async created() {
        this.loading = true;
        try {
            let roleRes = await getRoleList({uid:this.uid});
            this.formatList(roleRes.data);
        }catch (err){
            this.$confirm(err.message);
        }finally{
            this.loading = false;
        }
    },
    methods:{
        formatList(data) {
            this.loading = false;
            if(!data){
                return ;
            }
            this.dataList = data['data_list'] ? data['data_list'] : [];
            if (this.dataList.length <= 0) {
                return;
            }
            this.totalPage = data['total_page'];
            this.currentPage = data['current_page'];
            this.totalCount = data['total_count'];
        },
    },
}
</script>