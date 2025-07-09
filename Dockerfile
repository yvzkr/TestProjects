# Multi-stage build for ZplDesigner
FROM mcr.microsoft.com/dotnet/sdk:8.0 AS build

WORKDIR /src

# Copy solution and project files
COPY ZplDesigner.sln ./
COPY ZplDesigner.Library/ZplDesigner.Library.csproj ZplDesigner.Library/
COPY ZplDesigner.WebAPI/ZplDesigner.WebAPI.csproj ZplDesigner.WebAPI/

# Restore dependencies
RUN dotnet restore

# Copy source code
COPY . .

# Build the application
RUN dotnet build --no-restore

# Publish the application
FROM build AS publish
RUN dotnet publish ZplDesigner.WebAPI/ZplDesigner.WebAPI.csproj -c Release -o /app/publish

# Final stage
FROM mcr.microsoft.com/dotnet/aspnet:8.0 AS final

WORKDIR /app

# Copy published application
COPY --from=publish /app/publish .

# Expose port
EXPOSE 8080

# Set environment variables
ENV ASPNETCORE_URLS=http://+:8080
ENV ASPNETCORE_ENVIRONMENT=Production

# Start the application
ENTRYPOINT ["dotnet", "ZplDesigner.WebAPI.dll"] 