package tests

//func TestDBConfigSerialization(t *testing.T) {
//	logger := logz.NewLogger("TestDBConfig")
//
//	// 🔥 Criando uma configuração inicial simulada
//	dbConfig := NewDBConfigWithArgs(
//		"GodoTestDB",
//		"/srv/apps/projects/gdbase/tests/bkp/godo-test-config.yaml",
//		true,
//		logger,
//		false,
//	)
//
//	// 🔄 Tentando serializar pra YAML
//	serializedData, err := yaml.Marshal(dbConfig)
//	require.NoError(t, err, "Erro ao serializar DBConfig para YAML")
//
//	// 🚀 Escrevendo para arquivo temporário
//	filePath := "/srv/apps/projects/gdbase/tests/bkp/godo-test-config.yaml"
//	err = os.WriteFile(filePath, serializedData, 0644)
//	require.NoError(t, err, "Erro ao salvar arquivo YAML")
//
//	// 💡 Lendo de volta pra garantir consistência
//	readData, err := os.ReadFile(filePath)
//	require.NoError(t, err, "Erro ao ler arquivo YAML gerado")
//
//	fmt.Println("✅ Teste de serialização: Arquivo YAML gerado com sucesso!")
//	fmt.Println(string(readData))
//}
