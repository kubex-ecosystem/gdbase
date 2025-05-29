package tests

// type MockReadCloser struct{}

// func (m *MockReadCloser) Read(p []byte) (n int, err error) {
// 	return 0, io.EOF
// }

// func (m *MockReadCloser) Close() error {
// 	return nil
// }

// type MockDockerClient struct {
// 	m.Mock
// }

// func (m *MockDockerClient) ImagePull(ctx context.Context, image string, options i.PullOptions) (io.ReadCloser, error) {
// 	args := m.Called(ctx, image, options)
// 	return args.Get(0).(io.ReadCloser), args.Error(1)
// }

// func (m *MockDockerClient) ContainerList(ctx context.Context, options c.ListOptions) ([]c.Summary, error) {
// 	args := m.Called(ctx, options)
// 	return args.Get(0).([]c.Summary), args.Error(1)
// }

// func (m *MockDockerClient) ContainerCreate(ctx context.Context, config *c.Config, hostConfig *c.HostConfig, networkingConfig *n.NetworkingConfig, platform *o.Platform, containerName string) (c.CreateResponse, error) {
// 	args := m.Called(ctx, config, hostConfig, networkingConfig, platform, containerName)
// 	return args.Get(0).(c.CreateResponse), args.Error(1)
// }

// func (m *MockDockerClient) ContainerStart(ctx context.Context, containerID string, options c.StartOptions) error {
// 	args := m.Called(ctx, containerID, options)
// 	return args.Error(0)
// }

// func (m *MockDockerClient) VolumeCreate(ctx context.Context, options v.CreateOptions) (v.Volume, error) {
// 	args := m.Called(ctx, options)
// 	return args.Get(0).(v.Volume), args.Error(1)
// }

// func (m *MockDockerClient) VolumeList(ctx context.Context, options v.ListOptions) (v.ListResponse, error) {
// 	args := m.Called(ctx, options)
// 	return args.Get(0).(v.ListResponse), args.Error(1)
// }

// func TestStartContainerWithValidInputsStartsContainer(t *testing.T) {
// 	mockClient := new(MockDockerClient)
// 	dockerService := &s.DockerService{Cli: mockClient}

// 	mockClient.On("ContainerCreate", m.Anything, m.Anything, m.Anything, m.Anything, m.Anything, "test-service").Return(c.CreateResponse{ID: "12345"}, nil)
// 	mockClient.On("ContainerStart", m.Anything, "12345", m.Anything).Return(nil)
// 	mockClient.On("ImagePull", m.Anything, m.Anything, m.Anything).Return(&MockReadCloser{}, nil)

// 	err := dockerService.StartContainer("test-service", "test-image", []string{"ENV_VAR=value"}, nil, nil)
// 	r.NoError(t, err)
// 	mockClient.AssertExpectations(t)
// }

// func TestStartContainerWithErrorDuringCreationReturnsError(t *testing.T) {
// 	mockClient := new(MockDockerClient)
// 	dockerService := &s.DockerService{Cli: mockClient}

// 	mockClient.On("ContainerCreate", m.Anything, m.Anything, m.Anything, m.Anything, m.Anything, "test-service").Return(c.CreateResponse{}, errors.New("creation error"))
// 	mockClient.On("ImagePull", m.Anything, m.Anything, m.Anything).Return(&MockReadCloser{}, nil)

// 	err := dockerService.StartContainer("test-service", "test-image", []string{"ENV_VAR=value"}, nil, nil)
// 	r.Error(t, err)
// 	r.Contains(t, err.Error(), "creation error")
// 	mockClient.AssertExpectations(t)
// }

// func TestCreateVolumeWithValidInputsCreatesVolume(t *testing.T) {
// 	mockClient := new(MockDockerClient)
// 	dockerService := &s.DockerService{Cli: mockClient}

// 	mockClient.On("VolumeList", m.Anything, m.Anything).Return(v.ListResponse{Volumes: []*v.Volume{}}, nil)
// 	mockClient.On("VolumeCreate", m.Anything, m.Anything).Return(v.Volume{Name: "test-volume"}, nil)

// 	err := dockerService.CreateVolume("test-volume", "/path/to/device")
// 	r.NoError(t, err)
// 	mockClient.AssertExpectations(t)
// }

// func TestCreateVolumeWithExistingVolumeSkipsCreation(t *testing.T) {
// 	mockClient := new(MockDockerClient)
// 	dockerService := &s.DockerService{Cli: mockClient}

// 	mockClient.On("VolumeList", m.Anything, m.Anything).Return(v.ListResponse{Volumes: []*v.Volume{{Name: "test-volume"}}}, nil)

// 	err := dockerService.CreateVolume("test-volume", "/path/to/device")
// 	r.NoError(t, err)
// 	mockClient.AssertExpectations(t)
// }

// func TestCreateVolumeWithErrorDuringCreationReturnsError(t *testing.T) {
// 	mockClient := new(MockDockerClient)
// 	dockerService := &s.DockerService{Cli: mockClient}

// 	mockClient.On("VolumeList", m.Anything, m.Anything).Return(v.ListResponse{Volumes: []*v.Volume{}}, nil)
// 	mockClient.On("VolumeCreate", m.Anything, m.Anything).Return(v.Volume{}, errors.New("creation error"))

// 	err := dockerService.CreateVolume("test-volume", "/path/to/device")
// 	r.Error(t, err)
// 	r.Contains(t, err.Error(), "creation error")
// 	mockClient.AssertExpectations(t)
// }
