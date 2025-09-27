#!/bin/bash

echo "🤖 GDBASE Bot Models Test Suite"
echo "=================================="

# Check if we're in the right directory
if [ ! -f "go.mod" ]; then
    echo "❌ Error: Please run this script from the GDBASE root directory"
    exit 1
fi

echo "📊 Checking Go environment..."
go version
echo ""

echo "🔧 Running go mod tidy..."
go mod tidy
echo ""

echo "🏗️ Testing compilation..."
if go build -o /tmp/gdbase_test .; then
    echo "✅ Compilation successful"
    rm -f /tmp/gdbase_test
else
    echo "❌ Compilation failed"
    exit 1
fi
echo ""

echo "🧪 Running bot models tests..."
echo ""

echo "1️⃣ Running quick start test..."
if go test -v ./tests -run TestBotModelsQuickStart; then
    echo "✅ Quick start test passed"
else
    echo "❌ Quick start test failed"
fi
echo ""

echo "2️⃣ Running enums test..."
if go test -v ./tests -run TestEnumsAndConstants; then
    echo "✅ Enums test passed"
else
    echo "❌ Enums test failed"
fi
echo ""

echo "3️⃣ Running integration tests (may take longer)..."
if go test -v ./tests -run TestTelegramModelIntegration; then
    echo "✅ Telegram integration test passed"
else
    echo "❌ Telegram integration test failed"
fi
echo ""

if go test -v ./tests -run TestWhatsAppModelIntegration; then
    echo "✅ WhatsApp integration test passed"
else
    echo "❌ WhatsApp integration test failed"
fi
echo ""

if go test -v ./tests -run TestDiscordModelIntegration; then
    echo "✅ Discord integration test passed"
else
    echo "❌ Discord integration test failed"
fi
echo ""

if go test -v ./tests -run TestUnifiedMessagingIntegration; then
    echo "✅ Unified messaging test passed"
else
    echo "❌ Unified messaging test failed"
fi
echo ""

echo "4️⃣ Running validation tests..."
if go test -v ./tests -run TestModelValidations; then
    echo "✅ Model validation tests passed"
else
    echo "❌ Model validation tests failed"
fi
echo ""

echo "🏃‍♂️ Running example..."
if go run examples/bot_models_example.go; then
    echo "✅ Example ran successfully"
else
    echo "❌ Example failed to run"
fi
echo ""

echo "📈 Running performance tests (optional)..."
if go test -v ./tests -run TestPerformance -short; then
    echo "✅ Performance tests passed"
else
    echo "⚠️ Performance tests skipped or failed"
fi
echo ""

echo "🎉 Test suite completed!"
echo ""
echo "📋 Next steps:"
echo "   1. Check docs/BOT_MODELS_USAGE.md for usage guide"
echo "   2. Run 'go test ./tests -v' for all tests"
echo "   3. Start integrating with your actual database"
echo "   4. Implement webhook handlers"
echo ""
echo "🚀 Happy coding with the new bot models!"