<template>
    <div class="role-index">
        <v-container fill-height fluid class="align-start">
            <v-row class="mx-auto pt-6" width="90%">
                <v-col>
                    <v-row class="page-title mb-0">
                        角色列表
                        <v-spacer></v-spacer>
                        <v-text-field
                            v-model="kw"
                            append-icon="mdi-magnify"
                            label="按账号搜索"
                            @click:append="onSearch">
                        </v-text-field>
                    </v-row>
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
                        <template v-slot:item.qq="{ item }">
                            <v-chip v-if="item.qq==='GM_master'" color="primary">
                                GM
                            </v-chip>
                            <v-chip v-else>
                                普通
                            </v-chip>
                        </template>
                        <template v-slot:item.action="{ item }">
                            <v-btn color="primary" class="mr-2" small>
                                详情
                            </v-btn>
                            <v-btn @click="addRole(item)" color="info" class="mr-2" small>
                                添加角色
                            </v-btn>
                            <v-btn @click="onDelete(item)" color="error" small>
                                删除
                            </v-btn>
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
        <add-role :show="addRoleShow" :mid="mid" @cancel="onCancel"></add-role>
    </div>
</template>

<script>
import {getAccountList,deleteAccount} from "@/api/account";

import AddRole from "@/components/AddRole";
export default {
    components:{
        AddRole,
    },
    data() {
        return {
            loading:false,
            addRoleShow:false,
            mid:null,
            headers: [
                {
                    text: '人物UID',
                    align: 'start',
                    sortable: true,
                    value: 'UID',
                },
                {
                    text: '账号',
                    sortable: true,
                    value: 'accountname',
                },
                {text: '权限', value: 'qq'},
                {text: '操作', value: 'action'},
            ],
            dataList:[],
            pageSize: Number(process.env.VUE_APP_PAGE_SIZE),
            totalCount:0,
            currentPage:1,
            totalPage:1,
            kw:null,
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
        onSearch(){
            if(!this.kw){
                return false;
            }
            this.getList({kw:this.kw});
        },
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
            getAccountList(pageParams).then((res)=>{
                this.loading = false;
                this.formatList(res.data);
            }).catch((err)=>{
                this.loading = false;
                this.$store.dispatch('message/error',err.message);
            });
        },
        async onDelete(item){
            const res = await this.$confirm('确定要删除此账号吗？', { title: '系统提示' });
            if(res){
                this.loading = true;
                deleteAccount({id:item.UID}).then((res)=>{
                    this.loading = false;
                    this.$store.dispatch('message/success','删除成功');
                    this.getList();
                }).catch((err)=>{
                    this.loading = false;
                    this.$store.dispatch('message/error',err.message);
                });
            }
        },
        pageChange(v){
            this.getList({page:v});
        },
        addRole(item){
            this.addRoleShow = true;
            this.mid = item['UID'];
        },
        onCancel(){
            this.addRoleShow = false;
        }
    },
}
</script>