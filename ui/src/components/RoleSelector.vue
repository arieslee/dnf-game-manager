<template>
    <v-card>
        <v-card-title>
            选择角色
            <v-spacer></v-spacer>
            <v-text-field
                v-model="roleKw"
                append-icon="mdi-magnify"
                label="按角色名搜索"
                @click:append="onSearchRole"
                single-line
                hide-details
            ></v-text-field>
        </v-card-title>
        <v-divider></v-divider>
        <v-card-text style="height: 300px;">

            <v-data-table
                v-model="selectedRole"
                :headers="roleHeaders"
                :items="roleList"
                :loading="roleLoading"
                :single-select="true"
                item-key="charac_no"
                show-select
                class="elevation-1"
                :hide-default-footer="true"
                :disable-pagination="true"
                loading-text="数据载入中..."
                no-data-text="暂无数据"
            >
                <template slot="footer.page-text" slot-scope="item">
                    {{item.itemsLength}}中的{{item.pageStart}} - {{item.pageStop}}项
                </template>
            </v-data-table>
            <div class="text-center mt-3">
                <v-pagination
                    v-model="rolePage"
                    :length="roleTotalPage"
                    :total-visible="5"
                    @input="rolePageChange"
                ></v-pagination>
            </div>
        </v-card-text>
        <v-divider></v-divider>
        <v-card-actions>
            <v-btn
                color="blue darken-1"
                text
                @click="roleClose">
                关闭
            </v-btn>
            <v-btn
                color="blue darken-1"
                text
                @click="roleSelect">
                确定
            </v-btn>
        </v-card-actions>
    </v-card>
</template>

<script>
import {getRoleList} from "@/api/account";
export default {
    data(){
        return {
            roleKw:null,
            roleList:[],
            rolePageSize: 50,
            roleTotalCount:0,
            rolePage:1,
            roleTotalPage:1,
            roleLoading:false,
            role:{
                m_id:null,
                charac_no:null,
                charac_name:null,
            },
            roleHeaders:[
                {
                    text: '角色ID', value: 'm_id',sortable:false,
                },
                {
                    text: '角色名', value: 'charac_name',sortable:false,
                },
            ],
            selectedRole:[],
        }
    },
    created() {
        this.fetchRoleList();
    },
    methods:{
        onSearchRole(){
            if(!this.roleKw){
                return ;
            }
            this.roleList = [];
            this.fetchRoleList({kw:this.roleKw});
        },
        rolePageChange(v){
            this.fetchRoleList({page:v});
        },
        fetchRoleList(params){
            this.$store.dispatch('loading/show');
            this.roleLoading = true;
            let pageParams = {
                page:this.rolePage,
                page_size:this.rolePageSize,
            };
            if (params){
                pageParams = Object.assign(pageParams,params);
            }
            getRoleList(pageParams).then((res)=>{
                this.roleFormatList(res.data);
            }).catch((err)=>{
                this.$store.dispatch('message/error', '无法获取角色列表');
            }).finally(()=>{
                this.roleLoading = false
                this.$store.dispatch('loading/hide');
            });
        },
        roleFormatList(data) {
            this.loading = false;
            if(!data){
                return ;
            }
            this.roleList = data['data_list'] ? data['data_list'] : [];
            if (this.roleList.length <= 0) {
                return false;
            }
            this.roleTotalCount = data['total_count'];
            this.rolePage = data['current_page'];
            this.roleTotalPage = data['total_page'];
        },
        roleSelect(){
            if(!this.selectedRole || this.selectedRole.length <= 0){
                this.$store.dispatch('message/error', '请选择角色');
                return false;
            }
            this.role = Object.assign(this.role,this.selectedRole[0]);
            this.charac_no = this.role.charac_no;
            this.$emit('on-confirm',this.role);
        },
        roleClose(){
            this.$emit('on-close');
        }
    },
}
</script>