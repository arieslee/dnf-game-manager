<template>
    <v-dialog
        v-model="show"
        persistent
        max-width="600px"
    >
        <v-card>
            <v-card-title>
                <span class="text-h5">添加/修改角色</span>
            </v-card-title>
            <v-card-text>
                <v-container>
                    <v-text-field
                        label="角色名称*"
                        v-model="form.charac_name"
                        autocomplete="false"
                        required
                        :loading = "submitLoading"
                        :rules="[rules.required]"
                    ></v-text-field>
                    <v-text-field
                        v-model="form.cera"
                        label="初始点卷"
                        type="number"
                        :loading = "submitLoading"
                        :rules="[rules.required]"
                    ></v-text-field>
                    <v-text-field
                        v-model="form.coin"
                        label="初始金币"
                        type="number"
                        :loading = "submitLoading"
                        :rules="[rules.required]"
                        required
                    ></v-text-field>
                    <v-text-field
                        v-model="form.lev"
                        label="初始级别"
                        type="number"
                        min="1"
                        max="100"
                        :loading = "submitLoading"
                        :rules="[rules.required]"
                        required
                    ></v-text-field>
                    <v-select
                        clearable
                        v-model="form.job"
                        :items="selectJob"
                        :item-text="'label'"
                        :item-value="'value'"
                        label="职业*"
                        :loading = "submitLoading"
                        :rules="[rules.required]"
                        required
                    ></v-select>
                </v-container>
            </v-card-text>
            <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn
                    color="blue darken-1"
                    text
                    @click="onCancel"
                >
                    取消
                </v-btn>
                <v-btn
                    color="blue darken-1"
                    text
                    @click="onConfirm"
                >
                    确定
                </v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script>
import {JOB} from "@/config/job";
import {addRole} from "@/api/account";

const defualtForm = {
    m_id:0,
    charac_name:null,
    job:0,
    lev:1,
    coin:0,
    cera:0,
};
export default {
    props:{
        show:{
            type: Boolean,
            default: false,
        },
        mid:{
            type: Number,
            default: null,
        },
    },
    data(){
        return {
            selectJob:JOB,
            form:{...defualtForm},
            submitLoading:false,
            rules: {
                required: value => !!value || "不能为空."
            },
        };
    },
    methods:{
        onCancel(){
            this.$emit('cancel');
            this.form = {...defualtForm};
        },
        onConfirm(){
            let params = {...this.form};
            params.m_id = this.mid;
            console.log(this.mid);
            this.$store.dispatch('loading/show');
            addRole(params).then((res)=>{
                this.$store.dispatch('loading/hide');
                this.$store.dispatch('message/success', '角色添加成功');
            }).catch((err)=>{
                this.$store.dispatch('loading/hide');
                this.$store.dispatch('message/error', err.message);
            })
        },
    },
}
</script>