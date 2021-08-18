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
                    >
                        <template v-slot:item.job="{ item }">
                            <v-chip :color="item.job_data.color" :text-color="item.job_data.text_color">{{ item.job_data.label }}</v-chip>
                        </template>
                    </v-data-table>
                    <div class="text-center mt-3">
                        <v-pagination
                            v-model="currentPage"
                            :length="totalPage"
                            :total-visible="7"
                            @input="pageChange"
                        ></v-pagination>
                    </div>
                </v-col>
            </v-row>
        </v-container>
    </div>
</template>

<script>
import {getRoleList} from "@/api/account";
import {JOB} from "@/config/job";

export default {
    data() {
        return {
            loading:false,
            headers: [
                {
                    text: '人物UID',
                    align: 'start',
                    sortable: true,
                    value: 'm_id',
                },
                {
                    text: '角色ID',
                    sortable: true,
                    value: 'charac_no',
                },
                {text: '角色名', value: 'charac_name'},
                {text: '等级', value: 'lev'},
                {text: '职业', value: 'job'},
                {text: '点卷', value: 'cera'},
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
        this.getList();
    },
    methods:{
        formatList(data) {
            this.loading = false;
            if(!data){
                return ;
            }
            this.dataList = data['data_list'] ? data['data_list'] : [];
            if (this.dataList.length <= 0) {
                return false;
            }
            this.totalCount = data['total_count'];
            this.currentPage = data['current_page'];
            this.totalPage = data['total_page'];
            for(let i =0;i<this.dataList.length;i++){
                let item = {...this.dataList[i]};
                let job = this.jobFormat(item.job,item.grow_type);
                if(!job){
                    this.dataList[i].job_data = {
                        label:'未知',
                    };
                }else{
                    this.dataList[i].job_data = job;
                }
            }
        },
        getList(params){
            this.loading = true;
            let pageParams = {
                page:this.currentPage,
                page_size:this.pageSize,
            };
            if (params){
                pageParams = Object.assign(pageParams, params);
            }
            getRoleList(pageParams).then((res)=>{
                this.loading = false;
                this.formatList(res.data);
            }).catch((err)=>{
                this.loading = false;
                this.$store.dispatch('message/error',err.message);
            });
        },
        jobFormat(job,growType){
            let item = JOB.find((ele)=>{
                if(Number(ele.value)===Number(job)){
                    return ele;
                }
            });
            if (!item){
                return null;
            }
            if(typeof item.grow_type[growType]!=='undefined'){
                return item.grow_type[growType];
            }
            return item;
        },
        pageChange(v){
            this.getList({page:v});
        }
    },
}
</script>