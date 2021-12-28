class EndpointState
  def initialize
    @hash = {}
    @apis = []
  end

  def add(scheme)
    @hash[scheme.api] = [] if @hash[scheme.api].nil?
    @apis << scheme.api unless @apis.include?(scheme.api)
    scheme.endpoints.each do |endpoint|
      @hash[scheme.api] << endpoint
    end
  end

  def sort(api)
    @hash[api].sort_by(&enum_root)
  end
end
