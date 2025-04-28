package conf

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestAbstractEnvironment(t *testing.T) {

	os.Setenv("A_B_C", "aBc")
	e := &AbstractEnvironment{}
	osenv := e.GetSystemEnvironment()

	osenvjson, _ := json.MarshalIndent(osenv, "", "  ")
	fmt.Printf("osenv:%s\n", osenvjson)

	fmt.Printf("a.b.c:%s\n", osenv["a.b.c"])
}

func BenchmarkReplaceKey(b *testing.B) {
	// 测试不同长度的字符串
	b.Run("short", func(b *testing.B) {
		key := "a_b_c"
		for i := 0; i < b.N; i++ {
			replaceKey(key)
		}
	})

	b.Run("medium", func(b *testing.B) {
		key := "a_b_c_d_e_f_g_h_i_j"
		for i := 0; i < b.N; i++ {
			replaceKey(key)
		}
	})

	b.Run("long", func(b *testing.B) {
		key := "a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_q_r_s_t_u_v_w_x_y_z"
		for i := 0; i < b.N; i++ {
			replaceKey(key)
		}
	})

	// 测试不同数量的下划线
	b.Run("few_underscores", func(b *testing.B) {
		key := "a_b_c"
		for i := 0; i < b.N; i++ {
			replaceKey(key)
		}
	})

	b.Run("many_underscores", func(b *testing.B) {
		key := "a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_q_r_s_t_u_v_w_x_y_z"
		for i := 0; i < b.N; i++ {
			replaceKey(key)
		}
	})

	// 测试连续下划线
	b.Run("consecutive_underscores", func(b *testing.B) {
		key := "a__b___c____d"
		for i := 0; i < b.N; i++ {
			replaceKey(key)
		}
	})
}
